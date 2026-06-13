const qs = require("qs");

const options = {
  order: {
    name: "asc",
  },
  where: {
    name: {
      neq: "John Doe",
    },
  }
}
const url = new URL("http://localhost:8852/users");
url.search = qs.stringify(options, {
  skipNulls: true,
  encodeValuesOnly: true,
});
console.log(url.toString());

fetch(url.toString()).then(response => response.json()).then(data => {
  console.log(data);
});