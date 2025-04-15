package flag

import (
	"fmt"
	"os"
	"os/exec"
	"server/global"
	"time"
)

// SQLExport 导出 MySQL 数据
func SQLExport() error {
	mysql := global.Config.Mysql

	timer := time.Now().Format("20060102")
	sqlPath := fmt.Sprintf("mysql_%s.sql", timer)
	// 构造远程执行的命令
	remoteCmd := fmt.Sprintf(
		"docker exec mysql mysqldump -u%s -p%s %s",
		mysql.Username,
		mysql.Password,
		mysql.DBName,
	)

	// 构造本地执行的 ssh 命令
	cmd := exec.Command("ssh", "root@192.168.24.101", remoteCmd)

	outFile, err := os.Create(sqlPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd.Stdout = outFile
	return cmd.Run()

}
