import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Router } from '@angular/router';
import { ReqProto } from 'src/app/services/proto'
import { NzMessageService } from 'ng-zorro-antd';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  // public classList: string[] = ["图书", "文具", "生活用品", "电子产品", "化妆用品", "服装鞋包", "其它"];
  public classList: any;
  public freeGoodsList: any;
  public goodsList: any;
  public requestsList: any;
  constructor(
    private reqProto: ReqProto,
    public router: Router,
    private nzMessageService: NzMessageService,
    private http: HttpClient,
  ) { }

  ngOnInit() {
    this.http.get("/api/getGoodsType", httpOptions).subscribe((res: any) => {
      this.classList = res.data;
      console.log("货品类型", res);
    })
    this.switchClass(0);
  }
  switchClass(gt_id: number) {
    let url = "/api/getGoodsByType";
    let data = {
      gt_id: gt_id,
      category_id: 1
    }
    this.reqProto.action = "POST";
    this.reqProto.data = data;
    this.http.post(url, this.reqProto, httpOptions).subscribe((res: any) => {
      this.goodsList = res.data;
      console.log("商品", res);
    })
    this.reqProto.data.category_id = 2;
    this.http.post(url, this.reqProto, httpOptions).subscribe((res: any) => {
      this.freeGoodsList = res.data;
      console.log("免费商品", res);
    })
    this.reqProto.data.category_id = 3;
    this.http.post(url, this.reqProto, httpOptions).subscribe((res: any) => {
      this.requestsList = res.data;
      console.log("需求", res);
    })
    

  }
  logout() {
    let that = this
    var url = "/api/logout"
    that.http.get(url,httpOptions).subscribe(res => {
      console.log(res);
      that.router.navigate(["/login"])
    });
  }
}
const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
  })
};