module.exports = {
  outputDir: "./../public",
  chainWebpack: config => {
    config.resolve.alias.set("@components", __dirname + "/src/components");
    config.resolve.alias.set("@data", __dirname + "/src/data");
    config.resolve.alias.set("@pages", __dirname + "/src/pages");
    config.resolve.alias.set("@utils", __dirname + "/src/utils");
  }
};
