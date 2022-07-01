
module.exports = {
    chainWebpack: config => {
        let limit = 9999999999999999;
        config.module
            .rule('images')
            .test(/\.(png|gif|jpg)(\?.*)?$/i)
            .use('url-loader')
            .loader('url-loader')
            .tap(options => Object.assign(options, {limit: limit}));
        config.module
            .rule('fonts')
            .test(/\.(woff2?|eot|ttf|otf|svg)(\?.*)?$/i)
            .use('url-loader')
            .loader('url-loader')
            .options({
                limit: limit
            });
    },
    css: {
        extract: {
            filename: '[name].css',
            chunkFilename: '[name].css'
        },
        loaderOptions: {
            sass: {
                //sass-loader  v8-  use key name "data"
                //sass-loader  v8   use key name "prependData"
                //sass-loader  v10+ use key name "additionalData"
                additionalData: `@import "~@/assets/css/color.scss";`
            }
        }
    },
    configureWebpack: {
        output: {
            filename: '[name].js'
        },
        optimization: {
            splitChunks: false
        }
    },
    devServer: {
        disableHostCheck: true
    }
};
