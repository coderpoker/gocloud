package google

import (
	gce "github.com/cloudlibz/gocloud/compute/gce"
	googlecontainer "github.com/cloudlibz/gocloud/container/googlecontainer"
	bigtable "github.com/cloudlibz/gocloud/database/bigtable"
	googledns "github.com/cloudlibz/gocloud/dns/googledns"
	"github.com/cloudlibz/gocloud/gocloudinterface"
	googleloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/googleloadbalancer"
	googlecloudfunctions "github.com/cloudlibz/gocloud/serverless/googlecloudfunctions"
	googlestorage "github.com/cloudlibz/gocloud/storage/googlestorage"
)

// Google  struct represents Google Cloud provider.
type Google struct {
	gce.GCE
	googlestorage.GoogleStorage
	googleloadbalancer.Googleloadbalancer
	googlecontainer.Googlecontainer
	googledns.Googledns
	googlecloudfunctions.Googlecloudfunctions
	bigtable.Bigtable
}

func (*Google) Compute() gocloudinterface.Compute {
	return &gce.GCE{}
}

func (*Google) Storage() gocloudinterface.Storage {
	return &googlestorage.GoogleStorage{}
}

func (*Google) LoadBalancer() gocloudinterface.LoadBalancer {
	return &googleloadbalancer.Googleloadbalancer{}
}

func (*Google) Container() gocloudinterface.Container {
	return &googlecontainer.Googlecontainer{}
}

func (*Google) DNS() gocloudinterface.DNS {
	return &googledns.Googledns{}
}

func (*Google) Serverless() gocloudinterface.Serverless {
	return &googlecloudfunctions.Googlecloudfunctions{}
}

func (*Google) Database() gocloudinterface.Database {
	return &bigtable.Bigtable{}
}
