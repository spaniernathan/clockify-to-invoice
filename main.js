const puppeteer = require('puppeteer');    
(async() => {    
const browser = await puppeteer.launch();
const page = await browser.newPage();    

await page.goto(`file://${__dirname}/output/filled.html`);    
await page.pdf({
  path: 'invoice.pdf',
  format: 'A4', 
});    
await browser.close();    
})();
