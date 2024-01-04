package pkg

import "EdgeSys/apps/business/service"

var testService service.TestService

func Testpkg() {
	testService.Test()
}
