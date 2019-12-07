import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Router } from '@angular/router';
import { ReqProto } from 'src/app/services/proto'
import { NzMessageService } from 'ng-zorro-antd';
import { LocalStorageService } from '../../services/local-storage.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  // public classList: string[] = ["图书", "文具", "生活用品", "电子产品", "化妆用品", "服装鞋包", "其它"];
  public classList: any;
  public goodList: any;
  public priceHidden: boolean = false;
  public inputValue: string = "";
  public gt_id: number = 0;
  public category_id: number = -1;
  
  constructor(
    private reqProto: ReqProto,
    public router: Router,
    private nzMessageService: NzMessageService,
    private http: HttpClient,
    private lSData: LocalStorageService
  ) { }

  private userInfo: any;
  ngOnInit() {
    this.lSData.set("isLogin", "false");
    
    this.http.get("/api/getGoodsType", this.httpOptions).subscribe((res: any) => {
      this.classList = res.data;
      console.log("货品类型", res);
    })
    this.search();
  }
    // tab切换货品种类
  switchType(gt_id: number) {
    this.gt_id = gt_id;
    this.inputValue = "";
    this.search();
  }
  switchCategory(category_id: number) {
    this.category_id = category_id;
    this.search();
  }
    search() {
      this.reqProto.action = "POST";
      this.reqProto.data = {
        gt_id: this.gt_id,
        category_id: this.category_id,
        key: this.inputValue
      }
      console.log(this.reqProto);
      this.http.post("/api/getGoodsByType", this.reqProto, this.httpOptions).subscribe((res: any) => {
        this.goodList = res.data;
        console.log("res:")
        console.log(res);
      })
    }
    // 获取时间
    cutTime(time: string): string {
      if (time === undefined || time == null) {
        console.warn("cutTime()获取time为undefined或null")
      }
      return time.substr(0, 10); //+ ' ' + time.substr(11, 5);
    }
  toDetail(good: any) {
    this.lSData.remove("goodInfo");
    this.lSData.setObject("goodInfo", good);
    this.router.navigate(["/detail"]);
  }

}
