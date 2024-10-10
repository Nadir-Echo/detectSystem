<template>
  <div class="image-container">
    <div id="lineChart" :style="{ width: '600px', height: '600px' }" />
    <div id="myChart" :style="{ width: '600px', height: '600px' }" />
  </div>
</template>

<script>
import {getChart} from '@/api/defect-data'
export default {
  name: 'Hello',
  data() {
    return {
      msg: 'Welcome to kalacloud.com'
    }
  },
  mounted() {
    this.drawLine()
    this.renderChart()
  },
  methods: {
    renderChart() {
      const chart = this.$echarts.init(document.getElementById('myChart'))

      // 在这里配置你的饼状图数据和样式
      const options = {
        // 配置项...
        title: {
          text: 'Referer of a Website',
          subtext: 'Fake Data',
          left: 'center'
        },
        tooltip: {
          trigger: 'item'
        },
        legend: {
          orient: 'vertical',
          left: 'left'
        },
        series: [
          {
            name: 'Access From',
            type: 'pie',
            radius: '50%',
            data: [
              { value: 1048, name: 'Search Engine' },
              { value: 735, name: 'Direct' },
              { value: 580, name: 'Email' },
              { value: 484, name: 'Union Ads' },
              { value: 300, name: 'Video Ads' }
            ],
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ]
      }

      chart.setOption(options)
    },
    drawLine() {
      // 基于刚刚准备好的 DOM 容器，初始化 EChart 实例
      const lineChart = this.$echarts.init(document.getElementById('lineChart'))
      var option
      option = {
        title: {
          text: '折线图标题', // 设置折线图的标题
          left: 'center' // 设置标题居中显示
        },
        xAxis: {
          type: 'category',
          data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            data: [120, 200, 150, 80, 70, 110, 130],
            type: 'bar'
          }
        ]
      }
      // 绘制图表
      lineChart.setOption(option)
    }
  }
}
</script>

<style>
.image-container {
  display: flex;
}

#lineChart,
#myChart {
  flex: 1;
  margin-right: 10px; /* 可根据需要调整图像之间的间距 */
}
</style>

