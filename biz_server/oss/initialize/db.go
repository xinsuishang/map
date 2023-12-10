package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/go-sql-driver/mysql"
	"msp/biz_server/oss/config"
	"msp/biz_server/oss/internal/infra/mysql/model/ent"
	"msp/common/simpleutils"
)

// InitDB to init database
func InitDB() *ent.Client {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := simpleutils.MySqlDSNFormatUtil(c.User, c.Password, c.Host, c.Port, c.Name)
	entClient, err := ent.Open(
		"mysql",
		fmt.Sprintf(dsn),
		ent.Log(klog.Debug),
	)
	if err != nil {
		klog.Fatal(err)
	}

	//if err = entClient.Schema.Create(context.Background()); err != nil {
	//	klog.Fatalf("failed creating schema resources: %v", err)
	//}
	return entClient
}
