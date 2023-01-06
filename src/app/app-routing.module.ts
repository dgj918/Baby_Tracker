import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { InputHomeComponent } from './input-home/input-home.component';

const routes: Routes = [
  {path: '', component: HomeComponent},
  {path: 'input', component: InputHomeComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
