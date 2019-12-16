import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { HomeComponent } from './components/home/home.component';
import { RegisterComponent } from './components/register/register.component';
import { ChatComponent } from './components/chat/chat.component';
import { FavourComponent } from './components/favour/favour.component';
import { GoodsManageComponent } from './components/goods-manage/goods-manage.component';
import { UploadComponent } from './components/goods-manage/upload/upload.component';
import { ManageComponent } from './components/goods-manage/manage/manage.component';
import { GoodDetailComponent } from './components/good-detail/good-detail.component';
import { PersonalCenterComponent } from './components/personal-center/personal-center.component';
import { ManagePageComponent } from './components/manage-page/manage-page.component';


const routes: Routes = [
  { path: '', redirectTo: "login", pathMatch: 'full' },
  { path: "login", component: LoginComponent },
  { path: "home", component: HomeComponent },
  { path: "register", component: RegisterComponent },
  { path: "chat", component: ChatComponent },
  { path: "favour", component: FavourComponent },
  { path: "detail", component: GoodDetailComponent },
  {
    path: "goodsmanage", component: GoodsManageComponent,
    children: [
      { path: "", component: UploadComponent },
      { path: "upload", component: UploadComponent },
      { path: "manage", component: ManageComponent }
    ]
  },
  { path: "personalcenter", component: PersonalCenterComponent },
  { path: "managepage", component: ManagePageComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
