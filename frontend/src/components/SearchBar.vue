<script>
export default {
  data() {
    return {
      address: "",
    };
  },
};
</script>

<template>
  <section>
    <o-field label="Search">
      <o-input
        v-model="address"
        placeholder="123 Main St. Apt 456"
        type="search"
        icon="search"
        icon-clickable
        @icon-click="search"
      />
    </o-field>
  </section>
</template>

<style scoped>
section {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}
</style>
