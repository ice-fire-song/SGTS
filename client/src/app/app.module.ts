import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './components/home/home.component';
import { LoginComponent } from './components/login/login.component';
import { NgZorroAntdModule, NZ_I18N, en_US } from 'ng-zorro-antd';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { registerLocaleData } from '@angular/common';
import en from '@angular/common/locales/en';
import { ReactiveFormsModule } from '@angular/forms';
import { ReqProto } from 'src/app/services/proto';
import { RegisterComponent } from './components/register/register.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';
import { ChatComponent } from './components/chat/chat.component';
import { FavourComponent } from './components/favour/favour.component';
import { GoodsManageComponent } from './components/goods-manage/goods-manage.component';
import { UploadComponent } from './components/goods-manage/upload/upload.component';
import { ManageComponent } from './components/goods-manage/manage/manage.component';
import { GoodDetailComponent } from './components/good-detail/good-detail.component';

import { LocalStorageService } from './services/local-storage.service';

import { UserInfoServiceService } from './services/user-info-service.service';
import { NavigationComponent } from './components/navigation/navigation.component';
import { PersonalCenterComponent } from './components/personal-center/personal-center.component';
import { ManagePageComponent } from './components/manage-page/manage-page.component';

registerLocaleData(en);

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    RegisterComponent,
    ForgotPasswordComponent,
    ChatComponent,
    FavourComponent,
    GoodsManageComponent,
    UploadComponent,
    ManageComponent,
    GoodDetailComponent,
    NavigationComponent,
    PersonalCenterComponent,
    ManagePageComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgZorroAntdModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ReactiveFormsModule
  ],
  providers: [{ provide: NZ_I18N, useValue: en_US }, ReqProto, LocalStorageService,UserInfoServiceService],
  bootstrap: [AppComponent]
})
export class AppModule { }
