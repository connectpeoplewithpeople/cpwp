import { Component, OnInit } from '@angular/core';

import { environment } from './../../../environments/environment';

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.css']
})
export class FooterComponent implements OnInit {
  // 노출 여부
  isHide: boolean = false;
  baseUrl: string = environment.baseUrl;

  constructor() { }

  ngOnInit() {
  }

}
