const { app } = require("./app.js");
const { env, db } = require("./config");

const PORT = env.PORT;
const DB = env.DATABASE_URL;

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});
