// index.js
const express = require('express');
const session = require('express-session');
const passport = require('passport');
const GoogleStrategy = require('passport-google-oauth20').Strategy;

const CLIENT_ID = process.env.GOOGLE_CLIENT_ID;
const CLIENT_SECRET = process.env.GOOGLE_CLIENT_SECRET;
const BACKEND_URL = process.env.BACKEND_URL || 'http://localhost:8080/api/login';

passport.use(new GoogleStrategy({
  clientID: CLIENT_ID,
  clientSecret: CLIENT_SECRET,
  callbackURL: 'http://localhost:3000/auth/callback',
  passReqToCallback: true
}, (req, accessToken, refreshToken, params, profile, done) => {
  // params.id_token にJWTが入る
  req.idToken = params.id_token;
  return done(null, profile);
}));

const app = express();
app.use(session({ secret: 'dev', resave: false, saveUninitialized: true }));
app.use(passport.initialize());

app.get('/', (_req, res) => res.send('<a href="/auth/google">Google Login</a>'));

app.get('/auth/google',
  passport.authenticate('google', { scope: ['openid', 'email', 'profile'], session: false })
);

// コールバックでIDトークンを取得→8080へ送信

const axios = require('axios');

app.get('/auth/callback',
  passport.authenticate('google', { failureRedirect: '/', session: false }),
  async (req, res) => {
    const idToken = req.idToken;
    if (!idToken) return res.status(500).send('id_token not found');
    console.log('Google ID Token:', idToken);

    // Person登録用データ（シーダーに寄せる）
    // 下4桁をランダムにする関数
    function randomizeCoord(base) {
      const intPart = Math.floor(base);
      const fracPart = base - intPart;
      // 0.0001〜0.9999のランダム値
      const rand = Math.floor(Math.random() * 10000) / 10000;
      return intPart + rand;
    }

    const personData = {
      Latitude: randomizeCoord(35.1815),
      Longitude: randomizeCoord(136.9066),
      Emoji: "😅",
      Sign: "Z",
      Categories: [],
      Features: [],
      // ISO8601形式で送信（例: 2025-09-14T13:30:00+09:00）
      SightingTime: "2025-09-14T13:30:00+09:00",
      RegisterUUID: "a1a2b3c4-d5e6-7f89-0abc-def123456789",
      Gender: "女性",
      Clothing: "制服",
      Accessories: "バッグ",
      Vehicle: "自動車",
      Behavior: "暴力",
      Hairstyle: "パーマ"
    };

    try {
      const response = await axios.post(
        'http://localhost:8080/api/persons',
        personData,
        {
          headers: {
            'Authorization': `Bearer ${idToken}`,
            'Content-Type': 'application/json'
          }
        }
      );
      res.set('Content-Type', 'application/json').send({
        message: 'Person登録成功',
        apiResponse: response.data
      });
    } catch (err) {
      console.error('API error:', err?.response?.data || err.message);
      res.status(500).send({
        message: 'Person登録失敗',
        error: err?.response?.data || err.message
      });
    }
  }
);

app.listen(3000, () => console.log('http://localhost:3000'));
