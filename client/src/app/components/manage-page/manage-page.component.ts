import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Router } from '@angular/router';
import { ReqProto } from 'src/app/services/proto'
import { NzMessageService } from 'ng-zorro-antd';
import { LocalStorageService } from '../../services/local-storage.service';
@Component({
  selector: 'app-manage-page',
  templateUrl: './manage-page.component.html',
  styleUrls: ['./manage-page.component.css']
})
export class ManagePageComponent implements OnInit {
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  public classList: any;
  selectedFree = "全部类型";
  selectedGood = "全部类型";
  selectedRequest = "全部类型";
  public inputValue: string = "";
  public gt_id: number = 0;
  public category_id: number = -1;
  public listOfData: any;
  public userStatus: number;
  userInfo: any;
  uid: number;
  constructor(
    private reqProto: ReqProto,
    public router: Router,
    private nzMessageService: NzMessageService,
    private http: HttpClient,
    private lSData: LocalStorageService
  ) { 
    this.http.get("/api/login").subscribe((res: any) => {
      this.userStatus = res.data;
      if (res.data.status) {
        this.reqProto.data = {
          username: res.data.username
        }
        this.http.post("/api/getUserInfo", this.reqProto, this.httpOptions).subscribe((res: any) => {
          console.log(res);
          this.userInfo = res.data;
          this.uid = this.userInfo.uid;
        })
      } else {
        this.userInfo = {
          username: "",
          head_sculpture_path: "assets/images/head_img.jpg"
        }
      }
    })
  }

  ngOnInit() {
    this.http.get("/api/getGoodsType", this.httpOptions).subscribe((res: any) => {
      this.classList = res.data;
      console.log("货品类型", res);
    })
    this.search()
  }

  // 根据type_name获取gt_id
  getGtID(type_name: string): number {
    if (type_name == "全部类型") {
      return 0;
    }
    for (var x of this.classList) {
      if (x.type_name == type_name) {
        return x.gt_id;
      }
    }
    return 0;
  }
  // 选择不同类型
  switchType(category_id: number) {
    var gt_id: number = 0;
    console.log(this.selectedFree,this.selectedGood,this.selectedRequest)
    if (category_id == -1) {
      gt_id = 0;
    } else if (category_id == 0) {
      gt_id = this.getGtID(this.selectedFree);
    } else if (category_id == 1) {
      gt_id = this.getGtID(this.selectedGood);
    } else if (category_id == 2) {
      gt_id = this.getGtID(this.selectedRequest);
    }
    console.log(category_id, "--", gt_id);
    this.gt_id = gt_id;
    this.category_id = category_id;
    this.inputValue = "";
    this.search();
  }
  // switchCategory(category_id: number) {
  //   this.category_id = category_id;
  //   this.search();
  // }
  // 检索货品
  search() {
    this.reqProto.action = "POST";
    this.reqProto.data = {
      gt_id: this.gt_id,
      category_id: this.category_id,
      key: this.inputValue
    }
    console.log(this.reqProto);
    this.http.post("/api/getGoodsByType", this.reqProto, this.httpOptions).subscribe((res: any) => {
      this.listOfData = res.data;
      console.log("res:")
      console.log(res);
    })
  }

  // 操作
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
        console.log(good_status,"---")
        if (good_status == 0) {
          this.listOfData.splice(0, this.listOfData.length - 1);
        }
        this.nzMessageService.info(opt + '成功');
      } else {
        this.nzMessageService.info(opt + '失败');
      }
      console.log(res);
    })
  }

  // 通过编码计算状态名称
  getCategory(status: number):string {
    if (status == 0) {
      return "删除"
    } else if (status == 1) {
      return "上架"
    } else if (status == 2) {
      return "下架"
    }
  }
  // 时间字符串的截取
  cutTime(time: string): string {
    if (time === undefined || time == null) {
      console.warn("cutTime()获取time为undefined或null")
    }
    return time.substr(0, 10); //+ ' ' + time.substr(11, 5);
  }
}
