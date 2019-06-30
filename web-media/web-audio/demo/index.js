// 创建音频上下文
const AudioContext = window.AudioContext || window.webkitAudioContext
const audioContext = new AudioContext()

// 加载声音
const audioElement = document.querySelector('audio')
const track = audioContext.createMediaElementSource(audioElement)

const playButton = document.querySelector('button')

playButton.addEventListener('click', function() {
  // 检查上下文是否处于挂起状态，由于自动播放政策
  if (audioContext.state === 'suspended') {
    audioContext.resume()
  }

  if (this.dataset.playing === 'false') {
    audioElement.play()
    this.dataset.playing = 'true'
  } else if (this.dataset.playing === 'true') {
    audioElement.pause()
    this.dataset.playing = 'false'
  }
}, false)

audioElement.addEventListener('ended', () => {
  playButton.dataset.playing = 'false'
}, false)

const gainNode = audioContext.createGain()
const volumeControl = document.querySelector('#volume')
volumeControl.addEventListener('input', function() {
  gainNode.gain.value = this.value
}, false)

const pannerOptions = { pan: 0 }
const panner = new StereoPannerNode(audioContext, pannerOptions)
const pannerControl = document.querySelector('#panner')
pannerControl.addEventListener('input', function() {
  panner.pan.value = this.value
}, false)

// 连接目的地
track.connect(gainNode).connect(panner).connect(audioContext.destination)
