<template>
    <div>
      <div class="nav">
        <div class="navScroll">
          <span>
          <b>
            <p>苹果</p>
          </b>
          <div style="font-size: .2rem;line-height:20px">
            <p>产品ID：202101200132x</p>
            <p>生产商：深圳龙岗苹果联合生产有限公司</p>
          </div>
          <p>_</p>
          </span>
          <div class="doc">
          </div>
          
        </div>
        <!-- 我的账单 start -->
        <div class="zdList">
          <div class="same zdList-l">
            <div class="text">
              <h3>星级</h3>
              <p>四星级</p>
              <div>
                <img src="../../assets/img/star.png" class="image" style = "width: 90%;height:90%;padding: 0.04rem 0.0rem;">
              </div>

            </div>
          </div>
          
          <div class="same zdList-l">
            <div class="text">
              <h3>产地</h3>
              <p>采摘自</p>
              <b>新疆伊犁</b>
            </div>
          </div>

          <div class="same zdList-r">
            <div class="text">
              <h3>生产日期</h3>
              <p>x132批次</p>
              <b>2020/12/24</b>
            </div>
          </div>
        </div>
        <!-- 我的账单 end -->

        <!-- 花呗头条 start -->
        <div class="powerList">
          <div class="powerListLi">
            <div class="News">
              <div class="lNews"></div>
              <div class="rNews">
                <ul class="ul">
                  <li v-for="item in news"><i></i>{{item.title}}</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <!-- 花呗头条 end -->

        <!-- 产品数据折线图 -->
        <div class="linechart">
          <div class="text">
            <p>农产品培育数据</p>
          </div>
          <!-- <div class="data">
            <div class="ldata"></div>
              <div class="rdata">
                <ul class="ul">
                  <li v-for="item in news"><i></i>{{item.title}}</li>
                </ul>
              </div>
          </div> -->
          <line-chart :chart-data="datacollection"></line-chart>
        </div>
        <!-- 产品数据折线图 end -->

        <!-- 功能 start -->
        <div class="powerList">
          <!-- <div class="powerListLi">
            <div class="li" style="padding-top: .4rem">
              <i class="icon iconfont icon-gift" style="color: #5fb5e8;"></i>
              <p>分期好货</p>
            </div>
            <div class="li" style="padding-top: .18rem">
              <i class="icon iconfont icon-dianyingb" style="color: #86bf3c; font-size: .7rem"></i>
              <p style="position:relative; top: -.09rem">看电影</p>
            </div>
            <div class="li" style="padding-top: .4rem">
              <i class="icon iconfont icon-jiudian" style="color: #df8447;"></i>
              <p>住酒店</p>
            </div>
            <div class="li" style="padding-top: .4rem">
              <i class="icon iconfont icon-koubei1" style="color: #ff7a7a;"></i>
              <p>吃喝玩乐</p>
            </div>
          </div>
          <div class="powerListLi">
            <div class="li" style="padding-top: .3rem">
              <i class="icon iconfont icon-hanglixiang" style="color:#f69ca0;font-size: .6rem;"></i>
              <p style="position:relative; top: -.09rem">境外购物</p>
            </div>
            <div class="li" style="padding-top: .3rem">
              <i class="icon iconfont icon-qicheqianlian-" style="color:#ff0024;font-size: .6rem;"></i>
              <p style="position:relative; top: -.09rem">买车险</p>
            </div>
            <div class="li" style="padding-top: .3rem">
              <i class="icon iconfont icon-kefu" style="color:#46d4be;font-size: .6rem;"></i>
              <p style="position:relative; top: -.09rem">帮助</p>
            </div>
            <div class="li"></div>
          </div> -->
        </div>
        <!-- 功能列表 end -->

        <!-- 花呗福利 start -->
        <div class="powerList">
          <div class="Prize">
            <div class="span">Hi~ <b>中深创</b>粉丝来拿福利啦</div>
            <div class="content">
              <div class="lft bd">
                <span>我的抽奖机会</span>
                <p>3<b>次</b></p>
                <i class="icon iconfont icon-mjiantou"></i>
              </div>
              <div class="lft">
                <span>我的会员积分</span>
                <p>853<b>分</b></p>
                <i class="icon iconfont icon-mjiantou"></i>
              </div>
            </div>
          </div>
        </div>
        <!-- 花呗福利 end -->

        <!-- 广告 start -->
        <div class="powerList" v-for="item in infiniteDate">
          <div class="Ad">
            <!-- <div class="AdTop">
              <strong>{{item.img}}</strong>
            </div> -->
            <p>{{item.text}}</p>
          </div>
        </div>
        <!-- 广告 end -->

      </div>
    </div>
</template>

<script>
import LineChart from '../charts/LineChart.js'

export default {
  props: ['infiniteDate'],
  data () {
    return {
      tran: 0,
      timerObj: null,
      news: [
        {
          index: 1,
          title: '本季度top热销水果 -- 新疆阿克苏苹果'
        },
        {
          index: 2,
          title: '本季度最受好评产品 -- 东北黑土大米'
        },
        {
          index: 3,
          title: '下季度应季水果 -- 阳山水蜜桃，明日开抢！！！！！'
        }
      ],
      newsList: 0,
      timerOther: null,
      datacollection: null
    }
  },
  components: {
    LineChart
  },
  methods: {
    fillData () {
      this.datacollection = {
        labels: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 25, 27, 28, 29, 30],
        datasets: [
          {
            label: '土壤肥料浓度',
            pointBackgroundColor: 'rgba(75,192,192,1)',
            borderColor: 'rgba(75,192,192,1)',
            backgroundColor: 'rgba(75,192,9,0)',

            data: [20, 33, 24, 12, 24, 23, 15, 25, 23, 16, 28, 14, 18, 5, 28, 24, 25, 13, 14, 25, 24, 23, 14, 25, 15, 18, 19, 29, 14, 23]
          }, {
            label: '土壤氮含量',
            borderColor: '#218a3d',
            pointBackgroundColor: '#218a3d',
            backgroundColor: 'rgba(75,192,9,0)',
            type: 'line',
            data: [4, 12, 24, 23, 15, 25, 1, 2, 4, 2, 1, 4, 5, 22, 3, 8, 5, 16, 28, 14, 18, 5, 28, 24, 12, 24, 23, 15, 2, 8]
          }, {
            label: 'CO2浓度',
            borderColor: '#f87979',
            pointBackgroundColor: '#f87979',
            backgroundColor: 'rgba(75,192,9,0)',
            data: [1, 2, 4, 2, 1, 4, 5, 22, 3, 5, 7, 18, 23, 22, 12, 8, 4, 9, 5, 16, 5, 17, 20, 12, 13, 14, 23, 22, 7, 13, 16, 14, 24]
          }
        ]
      }
    },
    fixScroll () {
      let that = this
      let first = document.getElementsByClassName('first')
      let docWdt = document.getElementsByClassName('doc')[0].offsetWidth
      let tanA = 0
      for (let i = 0; i < first.length; i++) {
        tanA++
        first[i].style.top = Math.cos(20 * tanA * 0.017453293) * docWdt * 0.01 - 0.65 + 6 + 'rem' // ((docWdt - firstWdt)/2)   // - firstWdt/2
        first[i].style.left = Math.sin(20 * tanA * 0.017453293) * docWdt * 0.01 - 0.65 + 6 + 'rem'
      }
      this.timerObj = setInterval(function () {
        let covRadius = document.getElementsByClassName('doc')[0]
        that.tran++
//        console.log(1)
        covRadius.style.transform = 'rotate(' + -that.tran * 20 + 'deg)'
      }, 2000)
    },
    newsScroll () {
      let that = this
      let News = document.getElementsByClassName('News')[0]
      let oUl = News.getElementsByClassName('ul')[0]
      let aLiLen = oUl.getElementsByTagName('li').length
      let oneStep = 0.42

      this.timerOther = setInterval(function () {
        that.newsList++
        if (that.newsList > aLiLen - 1) {
          that.newsList = 0
        }
//        oUl.style.top = -that.newsList * oneStep + 'rem'
        oUl.style.top = -that.newsList * oneStep + 'rem'
      }, 2000)
    }
  },
  mounted () {  // created
    // this.transDiv() // 纯图片的旋转
    this.fixScroll()
    this.newsScroll()
    this.fillData()
  },
  beforeDestroy () {
    clearInterval(this.timerObj)
    clearInterval(this.timerOther)
  }
}
</script>

<style lang="less" scoped>
  .icon {
    width: 1rem; height: 1rem;
    vertical-align: -0.15rem;
    fill: currentColor;
    overflow: hidden;
  }
  .nav{
    padding-top: .96rem;
    width: 7.2rem;
    background-color: #fff;
    .navScroll{
      position: relative;
      width: inherit;
      height: 4.6rem;
      overflow: hidden;
      z-index: 10;
      background: url("../../assets/img/apple.jpg") no-repeat;
      background-size: 100%;
      .morning{
        position: absolute;
        left: 1.08rem;
        top: 1.17rem;
        width: 1.27rem;
        height: .4rem;
        z-index: 9;
        background: url("../../assets/img/morning.png") no-repeat;
        background-size: contain;
        animation: fade 2s infinite;
      }
      span{
        width: 5.8rem;
        position: absolute;
        bottom: -0.62rem;
        left: 0.36rem;
        z-index: 10;
        font-size: .4rem;
        color: rgba(3,3,3,0.8);
        background: rgba(254,254,254,0.9);
        border: 1px solid grey;
        padding: .1rem .3rem;
        border-radius: .2rem;
        line-height:40px;
      }
      .radiusDiv{
        width: 7.2rem;
        height: 7.2rem;
        border-radius: 50%;
        transform: scale(1.7);
        margin-top: 5.4rem;
        box-shadow: 0 0 .6rem rgba(16,143,233,0.4);
      }
      .covRadius{
        width: 7.2rem;
        height: 7.2rem;
        border-radius: 50%;
        transform: scale(1.1);
        background: url("../../assets/img/20x18.png") no-repeat;
        background-size: contain;
        transition: .3s;
      }
    }
    .zdList{
      display: flex;
      height: 2.3rem;
      border-top: 1px solid #e5e5e5;
      box-sizing:border-box;
      padding-bottom: .17rem;
      background-color: #f5f5f9;
      .zdList-l{
        flex:1;
        box-sizing: border-box;
        border-right: 1px solid #e5e5e5;
        background-color: #fff;
      }
      .zdList-r{
        flex:1;
        background-color: #fff;
      }
      .same{
        .text{
          padding: 0.4rem 0 0 0.6rem;
          h3{
            font-weight: normal;
            font-size: .3rem;
            color: #40454b;
          }
          p,b{
            color: #afb6c3;
          }
          p{
            padding-top: .16rem;
            padding-bottom: .02rem;
            font-size: .22rem;
          }
          b{
            font-weight: normal;
            font-size: .3rem;
          }
        }
      }
    }

    .linechart{
      padding-bottom: .2rem;
      // background-color: #f5f5f9;
      .text{
        padding-top: .3rem;
        font-size: 0.4rem;
        text-align: center;
        line-height: 50px;
        color: #464b51;
      }
    }

    .powerList{
      padding-bottom: .2rem;
      background-color: #f5f5f9;
      .powerListLi{
        display:flex;
        .News{
          display: flex;
          width: 7.2rem;
          background-color: #fff;
          padding: .25rem .2rem;
          .lNews{
            width: .62rem;
            height: .42rem;
            background: url("../../assets/img/apple.png") no-repeat;
            background-size: contain;
          }
          .rNews{
            position: relative;
            width: 5.5rem;
            height: .42rem;
            overflow: hidden;
            .ul{
              position: absolute;
              left:0;
              top: 0;
              transition: .3s;
              padding-left: .2rem;
              width: 5.5rem;
              height: .42rem;
              li{
                height: .42rem;
                line-height: .52rem;
                font-size: .2rem;
                color:#63676c;
                i{
                  display: inline-block;
                  margin-right: .1rem;
                  width: .1rem;
                  height: .1rem;
                  background-color: #ccc;
                  position: relative;
                  top: -.05rem;
                  border-radius: 50%;
                }
              }
            }
          }
        }
        .li{
          flex:1;
          height: 1.73rem;
          background-color: #fff;
          text-align:center;
          box-sizing:border-box;
          p{
            color: #464b51;
          }
          i{
            font-size: .5rem;
          }
        }
      }
    }
  }

  .doc{
    position: absolute;
    left: -2.397rem;
    top: 3.1rem;
    z-index: 999;
    width: 12rem;
    height: 12rem;
    transition: .3s;
  }
  .first{
    width: 1.3rem;
    height: 1.3rem;
    border-radius: 50%;
    position: absolute;
    /* 假如这个是45deg */
  }
  .a{
    background: url("../../assets/img/baozi.png");
    background-size: contain;
  }
  .b{
    background: url("../../assets/img/hbg.png");
    background-size: contain;
  }
  .c{
    background: url("../../assets/img/hlb.png");
    background-size: contain;
  }

  .Prize{
    background-color: #fff;
    .content{
      display: flex;
      padding: .3rem 0;
      .bd{
        border-right: 1px solid #eee;
      }
      .lft{
        text-indent: .5rem;
        box-sizing: border-box;
        height: 1.3rem;
        flex:1;
        position:relative;
        span{
          font-size: .3rem;
          color: #53585e;
        }
        p{
          font-size: .4rem;
          color: #9fa8b6;
          font-weight: normal;
          b{
            font-weight: normal;
            font-size: .15rem;
          }
        }
        i{
          text-indent: 0;
          position: absolute;
          right: -.3rem;
          top: .5rem;
          font-weight: bold;
          color: #7dc4e8;
        }
       }
    }
    .span{
      height: 1.2rem;
      line-height: 1.2rem;
      font-size: .2rem;
      color: #6a6a6a;
      text-indent: .18rem;
      border-bottom: 1px solid #eee;
    }
  }
  .Ad{
    .AdTop{
      img{
        display: block;
        width: 7.2rem;
        height: 3.45rem;
      }
      strong{
        display: block;
        width: 7.2rem;
        height: 2.45rem;
        line-height: 2.45rem;
        text-align: center;
        font-size: .3rem;
      }
    }
    p{
      background-color: #fff;
      height: .7rem;
      line-height: .7rem;
      font-size: .3rem;
      color: #1792ef;
      text-align: center;
    }
  }

  @keyframes trans {
    from{
      transform:scale(1.1) rotate(0);
    }
    to{
      transform:scale(1.1) rotate(360deg);
    }
  }
  @keyframes fade {
    0,100%{
      opacity:1;
    }
    50%{
      opacity:0;
    }
  }
</style>
