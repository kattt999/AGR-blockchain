<template>
  <div>
    <div class="image-wrapper">
      <div class="wrapper" style="width: 100%;height: 100%">
        <img src="src/assets/agri1.jpeg" draggable="false" style="display: block; height: auto;">
      </div>
    </div>

      <div class="movie-group">
        <img src="src/assets/bg.jpg" alt="">
        <div class='movie-desc-group'>
            <h2>食品溯源</h2>
        </div>
      </div>


    <div>
      <p style="line-height:0; margin-bottom:5px;">
                  <span data-boldtype="0"
                        style="font-family:SourceHanSansSC; font-weight:400; font-size:18px; color:rgb(16, 16, 16); font-style:normal; letter-spacing:0px; line-height:0px; text-decoration:none;">
                      rate: 4
                  </span>
      </p>
    </div>
    <div>
      <p style="line-height:0; margin-bottom:5px;">
                  <span data-boldtype="0"
                        style="font-family:SourceHanSansSC; font-weight:400; font-size:18px; color:#218a3d; font-style:normal; letter-spacing:0px; line-height:27px; text-decoration:none;">
                          绿色食品
                  </span>
      </p>
    </div>
    <div>
      <p style="line-height:0; margin-bottom:5px;">
                  <span data-boldtype="0"
                        style="font-family:SourceHanSansSC; font-weight:400; font-size:18px; color:#218a3d; font-style:normal; letter-spacing:0px; line-height:27px; text-decoration:none;">
                          有机食品
                  </span>
      </p>
    </div>
    <div>
      <p style="line-height:0; margin-bottom:5px;">
                  <span data-boldtype="0"
                        style="font-family:SourceHanSansSC; font-weight:400; font-size:18px; color:#218a3d; font-style:normal; letter-spacing:0px; line-height:27px; text-decoration:none;">
                          无公害食品
                  </span>
      </p>
    </div>

    <line-chart :chart-data="datacollection"></line-chart>
  </div>
</template>


<script>
  import LineChart from './LineChart.js'

  export default {
    components: {
      LineChart
    },
    data () {
      return {
        datacollection: null
      }
    },
    mounted () {
      this.fillData()
    },
    methods: {
      fillData () {
        this.datacollection = {
          labels: [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,25,27,28,29,30],
          datasets: [
            {
              label: 'fertilizer pooring',
              pointBackgroundColor: "rgba(75,192,192,1)",
              borderColor: "rgba(75,192,192,1)",
              backgroundColor: "rgba(75,192,9,0)",

              data: [20,33,24,12,24,23,15,25,23,16,28,14,18,5,28,24,25,13,14,25,24,23,14,25,15,18,19,29,14,23]
            }, {
              label: 'soil nutrition',
              borderColor: '#218a3d',
              pointBackgroundColor: "#218a3d",
              backgroundColor: "rgba(75,192,9,0)",
              type:"line",
              data: [4,12,24,23,15,25,1,2,4,2,1,4,5,22,3,8,5,16,28,14,18,5,28,24,12,24,23,15,2,8]
            }, {
              label: 'CO2',
              borderColor: '#f87979',
              pointBackgroundColor: "#f87979",
              backgroundColor: "rgba(75,192,9,0)",
              data: [1,2,4,2,1,4,5,22,3,5,7,18,23,22,12,8,4,9,5,16,5,17,20,12,13,14,23,22,7,13,16,14,24]
            }
          ]
        }
      },

    }
  }
</script>


<style>
  .small {
    max-width: 600px;
    margin:  150px auto;
  }
  .section{
      padding:0 1rem;
      box-sizing: border-box;
  }
</style>



