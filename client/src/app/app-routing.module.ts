import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { HomeComponent } from './components/home/home.component';
import { RegisterComponent } from './components/register/register.component';
import { ChatComponent } from './components/chat/chat.component';
import { FavourComponent } from './components/favour/favour.component';
import { GoodsManageComponent } from './components/goods-manage/goods-manage.component';


const routes: Routes = [
  { path: '', redirectTo: "home", pathMatch: 'full' },
  { path: "login", component: LoginComponent },
  { path: "home", component: HomeComponent },
  { path: "register", component: RegisterComponent },
  { path: "chat", component: ChatComponent },
  { path: "favour", component: FavourComponent },
  { path: "goods-manage", component: GoodsManageComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
