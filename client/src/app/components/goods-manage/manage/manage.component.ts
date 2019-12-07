import { Component, OnInit } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto } from 'src/app/services/proto'
import { LocalStorageService } from '../../../services/local-storage.service';
import { Router } from '@angular/router';
@Component({
  selector: 'app-manage',
  templateUrl: './manage.component.html',
  styleUrls: ['./manage.component.css']
})
export class ManageComponent implements OnInit {
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  public uid: number = 1;
  public goodList: any;
  public priceHidden: boolean = false;
  public inputValue: string = "";
  public category_id: number = 0;
  constructor(
    private nzMessageService: NzMessageService,
    public http: HttpClient,
    private reqProto: ReqProto,
    private router: Router,
    private lSData: LocalStorageService
  ) { }
 
  ngOnInit() {
    this.lSData.set("isLogin", "false");
    this.search();
  }
  // 跳转到货品详情
  toDetail(good: any) {
    // this.lSData.remove("goodInfo");
    // this.lSData.setObject("goodInfo", good);
    this.router.navigate(["/detail"]);
  }
  // tab切换货品种类
  switch(category_id: number) {
    this.inputValue = "";
    if (category_id == 1) {
      this.priceHidden = true;
    } else {
      this.priceHidden = false;
    }
    this.category_id = category_id;
    this.search();
  }
  search() {
    this.reqProto.action = "POST";
    this.reqProto.data = {
      uid: this.uid,
      category_id: this.category_id,
      key: this.inputValue
    }
    this.http.post("/api/getGoodsByCategory", this.reqProto, this.httpOptions).subscribe((res: any) => {
      this.goodList = res.data;
      console.log(res);
    })
  }
  cancel(): void {
    this.nzMessageService.info('取消成功');
  }
 // 0 删除、1 上架、2下架
  confirm(opt: string): void {
    let good_status: number;
    if (opt == "下架") {
      good_status = 2;
    } else if (opt == "删除") {
      good_status = 0;
    }
    this.reqProto.action = "POST";
    this.reqProto.data = {
      good_status: good_status,
    }
    this.http.post("/api/modifyGoodStatus", this.reqProto, this.httpOptions).subscribe((res: any) => {
      if (res.data.resultCode == 1) {
        this.goodList.splice(0, this.goodList.length - 1);
        this.nzMessageService.info(opt + '成功');
      } else {
        this.nzMessageService.info(opt + '失败');
      }
      console.log(res);
    })
  }
  // 获取时间
  cutTime(time: string): string {
    if (time === undefined || time == null) {
      console.warn("cutTime()获取time为undefined或null")
    }
    return time.substr(0, 10) + ' ' + time.substr(11, 5);
  }
}
