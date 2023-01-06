import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MilestoneInputComponent } from './milestone-input.component';

describe('MilestoneInputComponent', () => {
  let component: MilestoneInputComponent;
  let fixture: ComponentFixture<MilestoneInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MilestoneInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MilestoneInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
