import { TestBed } from '@angular/core/testing';

import { DatetimeConversionService } from './datetime-conversion.service';

describe('DatetimeConversionService', () => {
  let service: DatetimeConversionService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DatetimeConversionService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
