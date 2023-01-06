import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { DiaperInputComponent } from '../diaper-input/diaper-input.component';
import { EatingInputComponent } from '../eating-input/eating-input.component';
import { SleepInputComponent } from '../sleep-input/sleep-input.component';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  inputIconList: any[];

  constructor(public dialog: MatDialog) { 
    this.inputIconList = [
      {
        title: 'Feeding',
        icon: 'pediatrics',
         
      },
      {
        title: 'Diaper',
        icon: 'baby_changing_station'
      },
      {
        title: 'Sleep',
        icon: 'bedtime'
      },
    ];
  }

  ngOnInit(): void {
  }

  openDialogs = (title: string) => {
    switch(title) {
      case "Feeding":
        this.openFeedingInput()
        break;
      case "Diaper":
        this.openDiaperInput()
        break;
      case "Sleep":
        this.openSleepInput()
        break;
    }
  }

  openFeedingInput() {
    console.log("Clicked")
    const dialogConfig = new MatDialogConfig();
    dialogConfig.minWidth = '90%'
    const dialogRef = this.dialog.open(EatingInputComponent, dialogConfig);

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }

  openDiaperInput() {
    const dialogRef = this.dialog.open(DiaperInputComponent);

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }

  openSleepInput() {
    const dialogRef = this.dialog.open(SleepInputComponent);

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }


}
