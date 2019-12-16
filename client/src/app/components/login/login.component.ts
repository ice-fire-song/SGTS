import { Component, OnInit } from '@angular/core';
import { ReqProto } from 'src/app/services/proto'
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LocalStorageService } from '../../services/local-storage.service';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  validateForm: FormGroup;

  submitForm(): void {
    for (const i in this.validateForm.controls) {
     let a = this.validateForm.controls[i].markAsDirty();
      let b = this.validateForm.controls[i].updateValueAndValidity();
      console.log(a,b)
    }

    let username1 = this.validateForm.controls.userName["value"]
    let password1 = this.validateForm.controls.password["value"]
    let that = this
    if (username1 == "" || password1 == "") {
      that.nzMessageService.info("账号密码不为空")
      return
    }
    var url = "/api/login"
    console.log(this.reqProto)
    let formData = {
      username: username1,
      password: password1
    }
    this.reqProto.data = formData;
    that.http.post(url, this.reqProto, httpOptions).subscribe((res: any) => {
      if (res.data) {
        this.router.navigate(['/managepage']);
      } else {
        this.router.navigate(['/home']);
      }
     
    });
  }

  constructor(
    private fb: FormBuilder,
    private reqProto: ReqProto,
    public router: Router,
    private nzMessageService: NzMessageService,
    private http: HttpClient,
    private lSData: LocalStorageService
  ) { }

  ngOnInit(): void {
    this.lSData.set("isLogin", "true");
    this.validateForm = this.fb.group({
      userName: [null, [Validators.required]],
      password: [null, [Validators.required]],
      remember: [true]
    });
  }
}
const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
  })
};
