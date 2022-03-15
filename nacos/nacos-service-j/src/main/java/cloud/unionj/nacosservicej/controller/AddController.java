package cloud.unionj.nacosservicej.controller;

import cloud.unionj.nacosservicej.service.StatService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

/**
 * @author: created by wubin
 * com.treeyee.cloud.report.svc.controller
 * 2022/3/11
 */
@RestController
class AddController {
  @Autowired
  private StatService statService;

  @GetMapping(value = "/add")
  public Map<String, Object> add(int x, int y) {
    return statService.add(x, y);
  }
}
