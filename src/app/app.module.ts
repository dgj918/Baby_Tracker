import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { DiaperInputComponent } from './diaper-input/diaper-input.component';
import { SleepInputComponent } from './sleep-input/sleep-input.component';
import { MilestoneInputComponent } from './milestone-input/milestone-input.component';
import { EatingInputComponent } from './eating-input/eating-input.component';
import { InputHomeComponent } from './input-home/input-home.component';
import { LogComponent } from './log/log.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { DetailsComponent } from './details/details.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import { MatCardModule} from '@angular/material/card';
import { CurrentFeedingStatusComponent } from './current-feeding-status/current-feeding-status.component';
import { HttpClientModule } from '@angular/common/http';
import {MatDialogModule} from '@angular/material/dialog';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatFormFieldModule} from '@angular/material/form-field';
import { MatNativeDateModule } from '@angular/material/core';
import {MatInputModule} from '@angular/material/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {MatSelectModule} from '@angular/material/select';
@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    DiaperInputComponent,
    SleepInputComponent,
    MilestoneInputComponent,
    EatingInputComponent,
    InputHomeComponent,
    LogComponent,
    DashboardComponent,
    DetailsComponent,
    CurrentFeedingStatusComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatIconModule,
    MatDialogModule,
    MatButtonModule,
    MatCardModule,
    HttpClientModule,
    MatDatepickerModule,
    MatFormFieldModule,
    MatNativeDateModule,
    MatInputModule,
    MatSelectModule,
    ReactiveFormsModule,
    FormsModule,
  ],
  providers: [MatDatepickerModule],
  bootstrap: [AppComponent]
})
export class AppModule { }
