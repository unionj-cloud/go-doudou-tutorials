package cloud.unionj.nacosservicej.service;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;

import java.util.Map;

/**
 * @author: created by wubin
 * com.treeyee.cloud.report.svc.service
 * 2022/3/11
 */
@FeignClient(name = "statsvc", path = "/api")
@Service
public interface StatService {
  @PostMapping(value = "/add")
  Map<String, Object> add(@RequestParam("x") int x, @RequestParam("y") int y);
}
