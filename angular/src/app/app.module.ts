import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from "@angular/router";
import { HashLocationStrategy, Location, LocationStrategy } from '@angular/common';
import { SuiModule } from 'ng2-semantic-ui';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { IndexComponent } from './view/index/index.component';
import { NotfoundpageComponent } from './view/error/notfoundpage/notfoundpage.component';

const appRoutes: Routes = [
  { path: '', component: IndexComponent },
  { path: '**', component: NotfoundpageComponent }
];

@NgModule({
  declarations: [
    AppComponent,
    IndexComponent,
    NotfoundpageComponent
  ],
  imports: [
    RouterModule.forRoot(
      appRoutes, {}
    ),
    FormsModule,
    SuiModule,
    BrowserModule,
    HttpClientModule
  ],
  providers: [
    Location, { provide: LocationStrategy, useClass: HashLocationStrategy }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
