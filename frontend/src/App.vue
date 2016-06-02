<template>
  <div id="app">
    <header>
      <strong>Fuel</strong>Supply
    </header>
    <main>
      <create-form></create-form>
      <list></list>
      <total></total>
    </main>
  </div>
</template>

<script>
import CreateForm from './components/CreateForm';
import List from './components/List';
import Total from './components/Total';

export default {
  components: {
    CreateForm,
    List,
    Total,
  },
};
</script>

<style>

* {
  box-sizing: border-box;
}

html {
  height: 100%;
}

body {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}

#app {
  color: #2c3e50;
  background-color: #fff;
  border: 1px solid rgb(225, 237, 245);
  border-radius: 4px;
  margin-top: -100px;
  overflow: hidden;
  width: 1024px;
  font-family: Source Sans Pro, Helvetica, sans-serif;
  text-align: center;
}

#app a {
  color: #42b983;
  text-decoration: none;
}

header {
  font-size: 2rem;
  text-align: center;
  background-color: rgb(225, 237, 245);
  padding: 1rem;
  width: 100%;
}
</style>
