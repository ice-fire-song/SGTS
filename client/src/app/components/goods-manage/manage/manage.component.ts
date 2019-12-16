import { Component, OnInit } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto } from 'src/app/services/proto'
import { LocalStorageService } from '../../../services/local-storage.service';
import { Router } from '@angular/router';
import { UserInfoServiceService } from 'src/app/services/user-info-service.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-manage',
  templateUrl: './manage.component.html',
  styleUrls: ['./manage.component.css']
})
export class ManageComponent implements OnInit {
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  
  public goodList: any;
  public priceHidden: boolean = false;
  public inputValue: string = "";
  public category_id: number = 0;
  public uid: number;
  private userInfo: any;
  subscription: Subscription;
  constructor(
    private nzMessageService: NzMessageService,
    public http: HttpClient,
    private reqProto: ReqProto,
    private router: Router,
    private lSData: LocalStorageService,
    private uIService: UserInfoServiceService
  ) { 
    this.subscription = this.uIService.titleObservable.subscribe((src:any) => {
      console.log('得到的userInfo:' + src);
      this.userInfo = JSON.parse(src);
      console.log("userinfo:",this.userInfo)
      if (this.userInfo != null) {
        this.uid = this.userInfo.uid;
        this.lSData.set("isLogin", "false");
        this.search();
      } else {
        console.log("manage: 未接收到来自navigation的userInfo，则用户不处于登录状态或出错，将跳转到login"); 
        this.router.navigate(["/login"]);
      }
    });
  }
 
  ngOnInit() {

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
  confirm(opt: string,gid: number): void {
    console.log("opt", opt)
    let good_status: number;
    if (opt == "下架") {
      good_status = 2;
    } else if (opt == "删除") {
      good_status = 0;
    }
    console.log("opt", opt)
    this.reqProto.action = "POST";
    this.reqProto.data = {
      gid: gid,
      good_status: good_status,
    }
    this.http.post("/api/modifyGoodStatus", this.reqProto, this.httpOptions).subscribe((res: any) => {
      console.log(res)
      if (res.data.isSuccess) {
        if (good_status == 0) {
          this.goodList.splice(0, this.goodList.length - 1);
        }
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
