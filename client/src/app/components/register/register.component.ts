import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ValidationErrors, Validators } from '@angular/forms';
import { Observable, Observer } from 'rxjs';
import { ReqProto } from 'src/app/services/proto'
import { Router } from '@angular/router';
import axios from "axios";
import { NzMessageService } from 'ng-zorro-antd';
import {HttpClient, HttpHeaders} from "@angular/common/http";
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  
  validateForm: FormGroup;
  ngOnInit(){}
  submitForm(value: any): void {
    for (const key in this.validateForm.controls) {
      this.validateForm.controls[key].markAsDirty();
      this.validateForm.controls[key].updateValueAndValidity();
    }
    console.log(value);
    let that = this
    var url = "/api/register"
    console.log(this.reqProto)
    this.reqProto.data = value;
    that.http.post(url, this.reqProto, httpOptions).subscribe((res:any) => {
      console.log(res);
      if (res.data.is_user_exist) {
        this.nzMessageService.warning("该用户名已被使用，请修改用户名");
      } else {
        if (res.data.success) {
          this.nzMessageService.info("注册成功，请登录");
          this.router.navigate(['/login'])
        } else {
          this.nzMessageService.error("注册失败，请重新注册");
        }
      }
   
    });
  }

  resetForm(e: MouseEvent): void {
    e.preventDefault();
    this.validateForm.reset();
    for (const key in this.validateForm.controls) {
      this.validateForm.controls[key].markAsPristine();
      this.validateForm.controls[key].updateValueAndValidity();
    }
  }

  validateConfirmPassword(): void {
    setTimeout(() => this.validateForm.controls.confirm.updateValueAndValidity());
  }

  userNameAsyncValidator = (control: FormControl) =>
    new Observable((observer: Observer<ValidationErrors | null>) => {
      setTimeout(() => {
        if (control.value === 'JasonWood') {
          // you have to return `{error: true}` to mark it as an error event
          observer.next({ error: true, duplicated: true });
        } else {
          observer.next(null);
        }
        observer.complete();
      }, 1000);
    });

  confirmValidator = (control: FormControl): { [s: string]: boolean } => {
    if (!control.value) {
      return { error: true, required: true };
    } else if (control.value !== this.validateForm.controls.password.value) {
      return { confirm: true, error: true };
    }
    return {};
  };

  constructor(
    private fb: FormBuilder,
    private reqProto: ReqProto,
    public router: Router,
    private nzMessageService: NzMessageService,
    private http: HttpClient,
  ) {
    this.validateForm = this.fb.group({
      username: ['', [Validators.required], [this.userNameAsyncValidator]],
      email: ['', [Validators.email, Validators.required]],
      password: ['', [Validators.required]],
      confirm: ['', [this.confirmValidator]],
      comment: ['', [Validators.required]]
    });
  }
}
const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
  })
};