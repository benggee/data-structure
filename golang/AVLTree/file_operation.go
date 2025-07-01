package avltree

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// FileOperation 文件相关操作工具
// 提供读取文件内容并进行简单分词的功能

// ReadFile 读取指定文件的内容，并将其中包含的所有词语放进words切片中
// filename: 要读取的文件名
// words: 用于存储分词结果的切片指针
// 返回值: 是否成功读取文件
func ReadFile(filename string, words *[]string) bool {
	// 参数检查
	if filename == "" || words == nil {
		fmt.Println("filename is empty or words is nil")
		return false
	}

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Cannot open %s: %v\n", filename, err)
		return false
	}
	defer file.Close()

	// 创建scanner用于读取文件
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()
		// 对每一行进行分词处理
		tokenizeLine(line, words)
	}

	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return false
	}

	return true
}

// tokenizeLine 对一行文本进行分词处理
// 这是一个简单的分词实现，将连续的字母字符组合成单词
// line: 要分词的文本行
// words: 用于存储分词结果的切片指针
func tokenizeLine(line string, words *[]string) {
	// 将文本转换为小写
	line = strings.ToLower(line)

	// 寻找字符串中从start位置开始的第一个字母字符的位置
	start := firstCharacterIndex(line, 0)

	for i := start + 1; i <= len(line); {
		if i == len(line) || !unicode.IsLetter(rune(line[i])) {
			// 找到一个完整的单词
			if i > start {
				word := line[start:i]
				*words = append(*words, word)
			}
			start = firstCharacterIndex(line, i)
			i = start + 1
		} else {
			i++
		}
	}
}

// firstCharacterIndex 寻找字符串s中，从start位置开始的第一个字母字符的位置
// s: 要搜索的字符串
// start: 开始搜索的位置
// 返回值: 第一个字母字符的位置，如果没有找到则返回字符串长度
func firstCharacterIndex(s string, start int) int {
	for i := start; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			return i
		}
	}
	return len(s)
}

// ReadFileWithBuffer 使用缓冲读取大文件
// 这个版本适用于大文件，使用更大的缓冲区提高读取效率
func ReadFileWithBuffer(filename string, words *[]string) bool {
	// 参数检查
	if filename == "" || words == nil {
		fmt.Println("filename is empty or words is nil")
		return false
	}

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Cannot open %s: %v\n", filename, err)
		return false
	}
	defer file.Close()

	// 创建带缓冲的reader
	reader := bufio.NewReader(file)

	// 读取整个文件内容
	content, err := reader.ReadString(0) // 读取到文件末尾
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return false
	}

	// 对文件内容进行分词处理
	tokenizeContent(content, words)

	return true
}

// tokenizeContent 对整个文件内容进行分词处理
// content: 文件内容
// words: 用于存储分词结果的切片指针
func tokenizeContent(content string, words *[]string) {
	// 将内容转换为小写
	content = strings.ToLower(content)

	// 按行分割内容
	lines := strings.Split(content, "\n")

	// 对每一行进行分词
	for _, line := range lines {
		tokenizeLine(line, words)
	}
}

// ReadFileByChunks 分块读取大文件
// 这个版本适用于非常大的文件，通过分块读取来避免内存溢出
func ReadFileByChunks(filename string, words *[]string, chunkSize int) bool {
	// 参数检查
	if filename == "" || words == nil {
		fmt.Println("filename is empty or words is nil")
		return false
	}

	if chunkSize <= 0 {
		chunkSize = 4096 // 默认块大小
	}

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Cannot open %s: %v\n", filename, err)
		return false
	}
	defer file.Close()

	// 创建带缓冲的reader
	reader := bufio.NewReader(file)

	// 分块读取文件
	buffer := make([]byte, chunkSize)
	var remainingContent string

	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			// 将当前块的内容与剩余内容合并
			content := remainingContent + string(buffer[:n])

			// 找到最后一个完整的单词边界
			lastWordEnd := findLastWordBoundary(content)

			// 处理完整的单词
			if lastWordEnd > 0 {
				tokenizeContent(content[:lastWordEnd], words)
			}

			// 保存剩余的不完整内容
			remainingContent = content[lastWordEnd:]
		}

		if err != nil {
			break
		}
	}

	// 处理剩余的内容
	if remainingContent != "" {
		tokenizeContent(remainingContent, words)
	}

	return true
}

// findLastWordBoundary 找到字符串中最后一个单词的边界
// content: 要搜索的内容
// 返回值: 最后一个完整单词的结束位置
func findLastWordBoundary(content string) int {
	for i := len(content) - 1; i >= 0; i-- {
		if !unicode.IsLetter(rune(content[i])) {
			return i + 1
		}
	}
	return 0
}
