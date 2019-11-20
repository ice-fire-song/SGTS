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

  constructor(
    private reqProto: ReqProto,
    public router: Router,
    private nzMessageService: NzMessageService,
    private http: HttpClient,
  ) { }

  ngOnInit() {
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