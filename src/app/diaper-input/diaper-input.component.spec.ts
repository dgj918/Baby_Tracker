import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DiaperInputComponent } from './diaper-input.component';

describe('DiaperInputComponent', () => {
  let component: DiaperInputComponent;
  let fixture: ComponentFixture<DiaperInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DiaperInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DiaperInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
