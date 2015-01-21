package db_test

import (
	"github.com/coreos/go-etcd/etcd"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf-experimental/routing-api/db/etcdrunner"

	"testing"
)

func TestDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Suite")
}

var etcdRunner *etcdrunner.ETCDRunner
var etcdClient *etcd.Client

var _ = BeforeSuite(func() {
	etcdRunner = etcdrunner.NewETCDRunner(5001+config.GinkgoConfig.ParallelNode, 1)
})

var _ = AfterSuite(func() {
	etcdRunner.Stop()
})

var _ = BeforeEach(func() {
	etcdRunner.Stop()
	etcdRunner.Start()
	etcdClient = etcdRunner.Client()
})