import { Injectable } from '@angular/core';
import * as moment from 'moment';


@Injectable({
  providedIn: 'root'
})
export class DatetimeConversionService {

  constructor() { }

  convertToUtc = (localTime: moment.Moment ) => {
    return moment(localTime).utc().format('YYYY-MM-DD HH:mm:ss Z');
  }

  concateDateTime = (date: any, time: any): string => {
    return moment.utc(date + ' ' + time, 'DD/MM/YYYY HH:mm').utc().format('DD/MM/YYYY HH:mm');
  }
}
