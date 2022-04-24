<template>
  <div class="alert alert-danger">
    <h3>Unexpected Error Encountered Processing Request</h3>
    <p>{{ this.text }}</p>
  </div>
</template>
<script>
export default {
  props: {
    text: String,
  },
};
</script>