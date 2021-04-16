import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { InputFormComponent } from './input-form/input-form.component';
import { LoginButtonComponent } from './login-button/login-button.component';
import { HomeComponent} from './home/home.component'

const routes: Routes = [
  {path: 'inputforum', component:InputFormComponent },
  {path: 'login', component:LoginButtonComponent},
  {path: 'home', component:HomeComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
