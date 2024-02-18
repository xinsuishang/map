package initialize

import (
	"entgo.io/ent/dialect/sql"
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/go-sql-driver/mysql"
	"msp/biz_server/oss/config"
	"msp/biz_server/oss/internal/infra/mysql/model/ent"
	"msp/common/simpleutils"
	"time"
)

// InitDB to init database
func InitDB() *ent.Client {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := simpleutils.MySqlDSNFormatUtil(c.User, c.Password, c.Host, c.Port, c.Name)
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		klog.Fatalf("failed to db resources: %v", err)
		return nil
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	entClient := ent.NewClient(
		ent.Driver(drv),
		ent.Log(klog.Debug),
	)

	//if err = entClient.Schema.Create(context.Background()); err != nil {
	//	klog.Fatalf("failed creating schema resources: %v", err)
	//}
	return entClient
}
