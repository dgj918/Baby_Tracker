import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { environment } from 'src/environments/environment';
import { FeedingService } from '../services/feeding.service';
import * as moment from 'moment';
import { MatDialogRef } from '@angular/material/dialog';
import { DatetimeConversionService } from '../services/datetime-conversion.service';

@Component({
  selector: 'app-eating-input',
  templateUrl: './eating-input.component.html',
  styleUrls: ['./eating-input.component.scss']
})
export class EatingInputComponent implements OnInit {
  todayDate: Date
  feedingFormGroup: FormGroup
  formulas: any[];
  timeNow: any;

  constructor(private feedingService: FeedingService,
    public dialogRef: MatDialogRef<EatingInputComponent>,
    private datetimeConversionService: DatetimeConversionService) {
    this.todayDate = new Date();
    this.timeNow = moment().format('HH:mm:ss');
    console.log(this.timeNow)
    this.feedingFormGroup = new FormGroup({
      selectedDate: new FormControl(new Date()),
      Start: new FormControl(this.timeNow),
      End: new FormControl(this.timeNow),
      OuncesConsumed: new FormControl(environment.Default.ouncesConsumed),
      FormulaID: new FormControl(environment.Default.formulaID),
      BabyID: new FormControl(environment.Default.BabyID)
    })
    this.formulas = [];
  }

  ngOnInit(): void {
    this.formulas = environment.formulas;
    this.feedingFormGroup.valueChanges.subscribe((formData) => { 
      console.log(formData);
    })
  }

  submitFeeding = () => {
    var formVals = this.feedingFormGroup.value
    let currDate = formVals.selectedDate
    formVals.Start = this.formatDateTime(formVals.Start, currDate);
    formVals.End = this.formatDateTime(formVals.End, currDate)
    this.feedingService.postCurrentFeeding(this.feedingFormGroup.value).subscribe((data) =>{
      this.dialogRef.close();
    })
  }

  formatDateTime = (unformatedTime: string, unformatedDate: Date): string => {
    let startTime = unformatedTime
    let startSliceHour = startTime.slice(0,2)
    let startSliceMinute = startTime.slice(3,5)
    let momDate = moment([unformatedDate.getFullYear(), unformatedDate.getMonth(), unformatedDate.getDate(), startSliceHour, startSliceMinute])
    let utcMomDate = momDate.utc()
    return utcMomDate.format('YYYY-MM-DD HH:mm:ss')
  }

}
