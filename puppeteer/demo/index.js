const puppeteer = require('puppeteer');

(async () => {
  const browser = await puppeteer.launch({
    defaultViewport: {
      width: 375,
      height: 812,
      isMobile: true,
      hasTouch: true,
    }
  });
  const page = await browser.newPage();
  await page.goto('https://example.com');
  await page.screenshot({path: 'example.png'});

  await page.goto('https://h5.youzan.com/wscvis/knowledge/index?kdt_id=42095703&p=contentshow&alias=362ydsxrqv2nb&qr=paidcontent_362ydsxrqv2nb', {waitUtil: 'networkidle0'});
  await page.screenshot({path: 'youzan.png'});

  await browser.close();
})();