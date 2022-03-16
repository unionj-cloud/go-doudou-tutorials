package cloud.unionj.nacosservicej.controller;

import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author: created by wubin
 * com.treeyee.cloud.report.svc.controller
 * 2022/3/11
 */
@RestController
class EchoController {
  @RequestMapping(value = "/echo/{string}", method = RequestMethod.GET)
  public String echo(@PathVariable String string) {
    return "Hello Nacos Discovery " + string;
  }
}
