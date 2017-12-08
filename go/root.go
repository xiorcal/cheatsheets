// Copyright © 2017 xiorcal <xiorcal@protonmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	//"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var cfgFile string

//Daemon is true if the app must run as daemon
var Daemon bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "TODO APPNAME",
	Short: "TODO short desc",
	Long:  `TODO long desc`,

	Run: func(cmd *cobra.Command, args []string) {

		outputfile := viper.GetString("outputFile")

		log.SetOutput(&lumberjack.Logger{
			Filename:   viper.GetString("logfile"),
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})

		if Daemon {
			args := os.Args[1:]
			i := 0
			for ; i < len(args); i++ {
				if args[i] == "-d" || args[i] == "--daemon" {
					break
				}
			}
			args = append(args[:i], args[i+1:]...)
			cmd := exec.Command(os.Args[0], args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Start()
			fmt.Println("[PID]", cmd.Process.Pid)
			ioutil.WriteFile(viper.GetString("pidfile"), []byte(string(123)), 0644)
			os.Exit(0)
		} else {

			pid, err := ioutil.ReadFile(viper.GetString("pidfile"))
			if err != nil {
				// file does not exist, nothing to do
			} else {
				viper.Set("pid", pid)
				Daemon = true
				viper.Set("daemon", true)
				err = os.Remove(viper.GetString("pidfile"))

				if err != nil {
					Warn("pid file was found but cannot be deleted")
				}
			}
			for i := 0; ; i++ {
				log.Printf("printing line %v", i)
				if i%29 == 0 {
					Warn("OH! this line was a multiple of 29 (%v)", i)
				}
				if i > 1500 {
					break
				}
			}
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is APPNAME.json)")
	RootCmd.PersistentFlags().StringVar(&cfgFile, "logfile", "APPNAME.log", "log file (default is APPNAME.log)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.PersistentFlags().BoolVarP(&Daemon, "daemon", "d", false, "run as background app")
	//viper stuff
	viper.BindPFlag("daemon", RootCmd.PersistentFlags().Lookup("daemon"))
	viper.BindPFlag("logfile", RootCmd.PersistentFlags().Lookup("logfile"))
	viper.Set("pidfile", "APPNAME.pid")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find current directory.
		workingDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in workingDir directory with name "APPNAME" (without extension).
		viper.AddConfigPath(workingDir)
		viper.SetConfigName("APPNAME")
	}
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

//Warn will log a warning
func Warn(str string, values ...interface{}) {
	log.Printf("Warning : "+str, values...)

}
