package configs

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

type logFormatter struct { // 有关日志的操作都是基于这个类实现的
	logrus.TextFormatter // 继承TextFormatter
	// TextFormatter用于定义日志的一些功能，比方说它需不需要颜色，需不需要做那个时间戳的拆分，以及是不是需要打印堆栈这些东西

}

// 整个log.go只做了一件事，就是自定义了一个格式化工具logFormatter，作为logrus的Formatter

func (f *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 在logFormatter类上定义一个Format方法，接受一个logrus.Entry类型的参数，返回两个对象
	// 这个函数的参数entry是一个日志的结构体，里面存储了原始日志的各种信息，比如日志的级别，日志的消息，日志的时间等等
	// 我们想要的日志信息都格式符合我们自己的要求，所以需要对entry中的原始日志中的各种信息进行格式化才行
	// Format用于定义最终输出的日志的格式，也就是将entry中的原始日志信息格式化成我们想要的形式

	// 通过调用这个prettyCaller我们就可以知道正在执行的文件是哪个文件，位于的是哪个栈帧，当前栈帧对应于文中中的哪行代码
	prettyCaller := func(frame *runtime.Frame) string {
		// 定义一个函数，接受一个runtime.Frame类型的参数，返回一个字符串
		// runtime.Frame是一个结构体，其实就是当前执行的函数的栈帧，里面包含了函数的信息，比如函数的名字，函数的文件，
		// 函数的行号等等，这个函数的作用是把函数的信息拼接成一个字符串，方便我们查看函数的调用者是谁

		_, fileName := filepath.Split(frame.File) // 在 runtime.Frame 结构体中，File 字段通常包含当前正在执行的代码文件的完整路径。
		// filepath.Split 函数，传入 frame.File 作为参数。filepath.Split 函数会将路径字符串分割成路径和文件名两部分。
		// _, fileName := filepath.Split(frame.File)：使用空白标识符 _ 来忽略返回的路径部分，只接收文件名。
		// 这个文件名 fileName 随后可以被用于各种目的，例如在日志记录中包含文件名信息，或者用于其他需要文件名的操作。

		return fmt.Sprintf("%s:%d", fileName, frame.Line) // 最后，使用 fmt.Sprintf 函数将文件名和行号拼接成一个字符串，作为返回值返回。
		// runtime.Frame 结构体表示运行时栈的某一帧，它包含了当前执行函数的信息。frame.Line 字段表示这一帧所对应的源代码行号
		// fmt.Sprintf 函数是一个格式化字符串的函数，它的作用是将指定的占位符替换为相应的值，并返回一个格式化后的字符串。
		// fmt.Sprintf 函数被用来将文件名 fileName 和行号 frame.Line 格式化成一个字符串，
		// 其中 %s 代表一个字符串占位符，%d 代表一个十进制数字占位符。fmt.Sprintf 会用 fileName 替换 %s 的位置，用 frame.Line 替换 %d 的位置。
		// 最终，fmt.Sprintf 会返回一个诸如 main.go:12 的字符串，其中 main.go 是文件名，12 是行号。
	}

	// 日志写入的规则
	var b *bytes.Buffer
	if entry.Buffer != nil { // 如果entry.Buffer不为空，那么就使用entry.Buffer，entry_buffer有什么作用？
		// entry.Buffer是一个bytes.Buffer类型的数据，它是一个字节缓冲器，用于存储日志消息。
		// 当entry.Buffer不为空时，说明日志消息已经被写入了一个字节缓冲器中，此时可以直接使用这个缓冲器来写入日志消息。
		b = entry.Buffer
	} else { // 如果entry.Buffer为空，那么就使用bytes.Buffer
		// 如果entry.Buffer为空，那么就使用bytes.Buffer来创建一个新的字节缓冲器。
		// 这个字节缓冲器是一个空的缓冲器，它没有任何数据
		b = &bytes.Buffer{}
	}

	// 开始定义具体的日志写入规则 也就是不同日志内容的格式，对应视频中的41.24
	// 日志时间
	b.WriteString(fmt.Sprintf("[%s]", entry.Time.Format(f.TimestampFormat)))

	// 日志级别
	b.WriteString(fmt.Sprintf("%s", strings.ToUpper(entry.Level.String())))

	// 调用信息（只有当entry中包含了调用者信息才会写入到日志bytes.Buffer类型的b中）
	if entry.HasCaller() { // 如果entry.HasCaller()返回true，说明entry中包含了调用者的信息
		// 调用者信息
		b.WriteString(fmt.Sprintf(" [%s]", prettyCaller(entry.Caller)))
	}

	// 日志消息
	b.WriteString(fmt.Sprintf(" %s\n", entry.Message))

	return b.Bytes(), nil // 如果Format执行过程中未出错，将会返回b.Bytes()，也就是b的字节切片，以及表示没有错误的nil
}
