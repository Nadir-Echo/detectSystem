<template>
  <div class="image-container">
    <div id="lineChart" :style="{ width: '600px', height: '600px' }" />
    <div id="myChart" :style="{ width: '600px', height: '600px' }" />
  </div>
</template>

<script>
import { getChart } from '@/api/defect-data'

export default {
  name: 'Hello',
  data() {
    return {
      msg: 'Welcome to kalacloud.com',
      chartData: [] // Add chartData array to store the received data
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
        
    },
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
            data: this.chartData.map(item => ({
              value: item.value,
              name: item.name
            })),
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
      const option = {
        title: {
          text: '折线图标题', // 设置折线图的标题
          left: 'center' // 设置标题居中显示
        },
        xAxis: {
          type: 'category',
          data: this.chartData.map(item => item.name) // Use chartData to populate xAxis data
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            data: this.chartData.map(item => item.value), // Use chartData to populate series data
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
