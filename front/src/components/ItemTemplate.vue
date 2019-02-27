<template>
  <div class="text-left">
    <b>#{{item.id}}</b>
    <span>{{ item.original_name }}</span>
  </div>
</template>

<script>
export default {
  props: {
    item: { required: true },
    searchText: { required: true }
  }
}
</script>