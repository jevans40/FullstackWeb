import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginButtonComponent } from './login-button/login-button.component';
import { InputFormComponent } from './input-form/input-form.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { SearchBarComponent } from './search-bar/search-bar.component';
import { HomeComponent } from './home/home.component';
import { ArtistCardComponent } from './artist-card/artist-card.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginButtonComponent,
    InputFormComponent,
    SearchBarComponent,
    HomeComponent,
    ArtistCardComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    NgbModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
