package cmd

/*
* @Author: sottxiong
* @Date:   2019-07-07 16:28:34
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-08 00:55:35
 */

import (
  "bufio"
  "github.com/fatih/color"
  "os"
  "regexp"
  "strings"
)

type Question struct {
  Name  string
  Tip   string
  ReTip string
  Re    string
  //Do func()
}

type Questions struct {
  qs []Question
}

var (
  cyan       = color.New(color.FgCyan)
  green      = color.New(color.FgGreen)
  blue       = color.New(color.FgBlue)
  red        = color.New(color.FgRed)
  yellow     = color.New(color.FgYellow)
  boldCyan   = cyan.Add(color.Bold)
  boldGreen  = green.Add(color.Bold)
  boldYellow = yellow.Add(color.Bold)
  boldBlue   = blue.Add(color.Bold)
  boldRed    = red.Add(color.Bold)
  questions  = &Questions{}
  answers    = map[string]string{}
)

func AddQuestion(name, tip, retip, re string) *Questions {
  questions.qs = append(questions.qs, Question{name, tip, retip, re})
  return questions
}

func ask_question(q string, color *color.Color) string {
  inputReader := bufio.NewReader(os.Stdin)

  color.Printf(q)
  inputData, err := inputReader.ReadString('\n')
  if err != nil {
    panic(err)
  }
  return strings.Trim(inputData, "\n")

}

func command(q *Question) {
  data := ask_question(q.Tip, boldGreen)
  re := regexp.MustCompile(q.Re)
  for {
    match := re.FindString(data)
    if len(match) > 0 {
      answers[q.Name] = data
      break
    } else {
      data = ask_question(q.ReTip, boldRed)
    }
  }
  //q.Do()
}

func Exec() map[string]string {
  for _, question := range questions.qs {
    command(&question)
  }
  return answers
}
