import { Component, OnInit } from '@angular/core';
import { FeedingService } from '../services/feeding.service';

@Component({
  selector: 'app-current-feeding-status',
  templateUrl: './current-feeding-status.component.html',
  styleUrls: ['./current-feeding-status.component.scss']
})
export class CurrentFeedingStatusComponent implements OnInit {
  feedingHistory: any[];

  constructor(private feedingSerivice: FeedingService) {
    this.feedingHistory = [];
  }

  ngOnInit(): void {
    this.getCurrentFeedingStatus();
  }

  getCurrentFeedingStatus = () => {
    this.feedingSerivice.getCurrentFeedingSubsetStatus(3).subscribe((data) => {
      console.log(data)
      this.feedingHistory = data.data;
    })
  }

}
