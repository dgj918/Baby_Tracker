import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';


@Injectable({
  providedIn: 'root'
})
export class FeedingService {

  constructor(
    private http: HttpClient
  ) { 
    
  }

  getCurrentFeedingStatus(): Observable<any> {
    return this.http.get(`${environment.apiUrl}/feeding`)
  }

  getCurrentFeedingSubsetStatus(num: number): Observable<any> {
    console.log(environment.apiUrl)
    // return this.http.get(`${environment.apiUrl}/feeding/subset/${num}`)
    return this.http.get(`http://192.168.68.70:8080/feeding/subset/${num}`)
  }

  postCurrentFeeding(feedingData: any): Observable<any> {
    return this.http.post(`${environment.apiUrl}/feeding`, feedingData);
  }
}
