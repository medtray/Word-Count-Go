package main

import (
  "testing"
  "io/ioutil"
  "bytes"
  "os/exec"
)

type files struct {
    in, out, want string
}

func TestWordCount_UNIX_1(t *testing.T) {
  c := files{"THEGODFATHER.dat", "THEGODFATHER_UNIX_WC.out", "THEGODFATHER_CORRECT_WC.dat"}
  WordCount_UNIX(c.in, c.out)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_UNIX_1(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_GO_1(t *testing.T) {
  c := files{"THEGODFATHER.dat", "THEGODFATHER_GO_WC.out", "THEGODFATHER_CORRECT_WC.dat"}
  WordCount_GO(c.in, c.out)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_GO_1(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_MR_S_1_M4_R4(t *testing.T) {
  c := files{"THEGODFATHER.dat", "THEGODFATHER_MR_S_1_M4_R4_WC.out", "THEGODFATHER_CORRECT_WC.dat"}
  WordCount_MR_S(c.in, c.out, 4, 4)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_S_1_M4_R4(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_MR_SMP_1_M4_R4(t *testing.T) {
  c := files{"THEGODFATHER.dat", "THEGODFATHER_MR_SMP_1_M4_R4_WC.out", "THEGODFATHER_CORRECT_WC.dat"}
  WordCount_MR_SMP(c.in, c.out, 4, 4)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_SMP_1_M4_R4(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_MR_SMP_1_M8_R8(t *testing.T) {
  c := files{"THEGODFATHER.dat", "THEGODFATHER_MR_SMP_1_M8_R8_WC.out", "THEGODFATHER_CORRECT_WC.dat"}
  WordCount_MR_SMP(c.in, c.out, 8, 8)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_SMP_1_M8_R8(%q, %q)", c.in, c.out)
  }
}

/*
//BONUS
func TestWordCount_MR_DMP_1_M8_R8(t *testing.T) {
  c := files{"THEGODFATHER.dat", "THEGODFATHER_MR_DMP_1_M8_R8_WC.out", "THEGODFATHER_CORRECT_WC.dat"}
  WordCount_MR_DMP(c.in, c.out, 8, 8)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_DMP_1_M8_R8(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_MR_DMP_1_M16_R16(t *testing.T) {
  c := files{"THEGODFATHER.dat", "THEGODFATHER_MR_DMP_1_M16_R16_WC.out", "THEGODFATHER_CORRECT_WC.dat"}
  WordCount_MR_DMP(c.in, c.out, 16, 16)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_DMP_1_M16_R16(%q, %q)", c.in, c.out)
  }
}
*/

func TestWordCount_UNIX_1000(t *testing.T) {
  c := files{"THEGODFATHER1000.dat", "THEGODFATHER1000_UNIX_WC.out", "THEGODFATHER1000_CORRECT_WC.dat"}
  WordCount_UNIX(c.in, c.out)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_UNIX_1000(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_GO_1000(t *testing.T) {
  c := files{"THEGODFATHER1000.dat", "THEGODFATHER1000_GO_WC.out", "THEGODFATHER1000_CORRECT_WC.dat"}
  WordCount_GO(c.in, c.out)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_GO_1000(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_MR_S_1000_M4_R4(t *testing.T) {
  c := files{"THEGODFATHER1000.dat", "THEGODFATHER_MR_S_1000_M4_R4_WC.out", "THEGODFATHER1000_CORRECT_WC.dat"}
  WordCount_MR_S(c.in, c.out, 4, 4)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_S_1000_M4_R4(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_MR_SMP_1000_M4_R4(t *testing.T) {
  c := files{"THEGODFATHER1000.dat", "THEGODFATHER_MR_SMP_1000_M4_R4_WC.out", "THEGODFATHER1000_CORRECT_WC.dat"}
  WordCount_MR_SMP(c.in, c.out, 4, 4)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_SMP_1000_M4_R4(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_MR_SMP_1000_M8_R8(t *testing.T) {
  c := files{"THEGODFATHER1000.dat", "THEGODFATHER_MR_SMP_1000_M8_R8_WC.out", "THEGODFATHER1000_CORRECT_WC.dat"}
  WordCount_MR_SMP(c.in, c.out, 8, 8)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_SMP_1000_M8_R8(%q, %q)", c.in, c.out)
  }
}

/*
//BONUS
func TestWordCount_MR_DMP_1000_M8_R8(t *testing.T) {
  c := files{"THEGODFATHER1000.dat", "THEGODFATHER_MR_DMP_1000_M8_R8_WC.out", "THEGODFATHER1000_CORRECT_WC.dat"}
  WordCount_MR_DMP(c.in, c.out, 8, 8)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_DMP_1000_M8_R8(%q, %q)", c.in, c.out)
  }
}

func TestWordCount_MR_DMP_1000_M16_R16(t *testing.T) {
  c := files{"THEGODFATHER1000.dat", "THEGODFATHER_MR_DMP_1000_M16_R16_WC.out", "THEGODFATHER1000_CORRECT_WC.dat"}
  WordCount_MR_DMP(c.in, c.out, 16, 16)
  f1, _ := ioutil.ReadFile(c.out)
  f2, _ := ioutil.ReadFile(c.want)
  if !bytes.Equal(f1, f2) {
    t.Errorf("WordCount_MR_DMP_1000_M16_R16(%q, %q)", c.in, c.out)
  }
}
*/

func TestCleanFiles(t *testing.T) {
  cmd := "rm THEGODFATHER*.out"
  exec.Command("bash","-c",cmd).Output()
}
