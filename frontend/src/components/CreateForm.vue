<template>
  <section>
    <form @submit.prevent="addFuelSupply">
      <input type="text" name="location" placeholder="Descrição" :model="newItem.location">
      <input type="text" name="quantity" placeholder="Quantidade" :model="newItem.quantity">
      <input type="text" name="value" placeholder="Valor" :model="newItem.value">
      <input type="text" name="km" placeholder="KM" :model="newItem.km">
      <input type="date" name="date" placeholder="Data" :model="newItem.date">
    </form>
  </section>
</template>

<script>
export default {
  data() {
    return {};
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
section {
  border-bottom: 2px solid rgb(225, 237, 245);
  padding: 1rem;
}

input {
  display: block;
  float: left;
  background-color: #fff;
  border: 1px solid #e2e2e2;
  border-radius: 4px;
  font-size: 0.9rem;
  padding: 8px;
  margin: 0 4px;
}

input[type="date"] {
  padding-bottom: 6px;
}

form {
  display: block;
  margin: 0 auto;
}

form::after {
  content: ' ';
  display: block;
  clear: both;
  float: none;
  width: 1px;
}

</style>
