import * as firebase from "firebase";
let firebaseConfig = {
  apiKey: "AIzaSyBEA1qHJu5hdKf2X-fsVpA_R91CPc-L-IE",
  authDomain: "stocksmith-1a671.firebaseapp.com",
  databaseURL: "https://stocksmith-1a671.firebaseio.com",
  projectId: "stocksmith-1a671",
  storageBucket: "stocksmith-1a671.appspot.com",
  messagingSenderId: "5580128612",
  appId: "1:5580128612:web:88cfa985f23331bd39ba71"
};
// Initialize Firebase
let fire = firebase.initializeApp(firebaseConfig);

export default fire;
