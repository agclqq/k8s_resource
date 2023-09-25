package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Node struct {
	Key        string
	Type       string
	FType      string
	Required   bool
	Annotation []string
	Children   []*Node
}

func Explain(res string) *Node {
	head := &Node{
		Key:      res,
		Children: make([]*Node, 0),
	}
	iterate(head)
	return head
}
func iterate(head *Node) {
	getOneExplain(head)
	//if head.Key == "job.spec" {
	//	return
	//}
	for _, v := range head.Children {
		iterate(v)
	}
}
func getOneExplain(fNode *Node) {
	childNode := &Node{Children: make([]*Node, 0)}
	isAnalyse := false
	_, err := Cmd("kubectl", []string{"explain", fNode.Key}, func(msg []byte) {
		strMsg := string(msg)
		if strings.HasPrefix(strMsg, "FIELDS:") {
			isAnalyse = true
			return
		}
		if !isAnalyse {
			return
		}
		if strings.HasPrefix(strMsg, "     ") {
			childNode.Annotation = append(childNode.Annotation, strings.TrimSpace(strMsg))
		} else if strings.HasPrefix(strMsg, "   ") {
			keyInfo := strings.Split(strings.TrimSpace(strMsg), "\t")
			required := false
			if len(keyInfo) < 2 {
				return
			}
			childNode = &Node{Children: make([]*Node, 0)}
			if len(keyInfo) > 2 && strings.Contains(keyInfo[2], "-required-") {
				required = true
			}
			childNode.Key = fNode.Key + "." + keyInfo[0]
			childNode.Type = keyInfo[1]
			childNode.FType = fNode.Type
			childNode.Required = required
			fNode.Children = append(fNode.Children, childNode)
		}
	})
	if err != nil {
		return
	}
}

func Cmd(commandName string, params []string, callback func(msg []byte)) (*os.ProcessState, error) {
	cmd := exec.Command(commandName, params...)
	fmt.Println("Cmd", commandName, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	cmd.Stderr = os.Stderr
	//cmd.Dir = dir
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的

	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil || io.EOF == err {
			break
		}
		callback(line)
	}
	err = cmd.Wait()

	return cmd.ProcessState, err
}
func (n *Node) ToYaml() {
	PrintToYaml(n, 0, nil)
}
func (n *Node) ToYamlFile(filePath ...string) error {
	fp := n.Key + ".yaml"
	if len(filePath) > 0 {
		fp = filePath[0]
	}
	file, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	PrintToYaml(n, 0, file)
	return nil
}
func PrintToYaml(head *Node, i int, file *os.File) {
	sj := strings.Repeat("  ", i)
	arrayLeading := ""
	arrayLeader := true
	for _, v := range head.Children {
		keys := strings.Split(v.Key, ".")
		if strings.Contains(v.FType, "[]") {
			if arrayLeader {
				arrayLeading = "- "
				arrayLeader = false
			} else {
				arrayLeading = "  "
			}
		}

		placeholder := "placeholder"
		if v.Type == "<integer>" {
			placeholder = "0"
		}
		if v.Type == "<boolean>" {
			placeholder = "false"
		}
		if strings.Contains(v.Type, "Object") || strings.Contains(v.Type, "map[string]") || strings.Contains(v.Type, "[]string") {
			placeholder = ""
		}
		row := fmt.Sprintf("%s%s%s: %s\t#%s\t%v\n", sj, arrayLeading, keys[len(keys)-1], placeholder, v.Type, v.Required) //空格缩进，数组前导符，打印key，值占位符，类型，是否必须
		if file != nil {
			file.WriteString(row)
		}
		fmt.Printf("%s", row)
		if strings.Contains(v.Type, "map[string]") {
			row = fmt.Sprintf("%s    %s: %s\n", sj, "placeholder", "placeholder")
			if file != nil {
				file.WriteString(row)
			}
			fmt.Printf("%s", row)
		}
		if strings.Contains(v.Type, "[]string") {
			row = fmt.Sprintf("%s    - %s\n", sj, "placeholder")
			if file != nil {
				file.WriteString(row)
			}
			fmt.Printf("%s", row)
		}
		for _, vv := range v.Annotation {
			row = fmt.Sprintf("%s\t#%s\n", sj, vv)
			if file != nil {
				file.WriteString(row)
			}
			fmt.Printf("%s", row)
		}

		if len(v.Children) > 0 {
			PrintToYaml(v, i+2, file)
		}
	}
}
