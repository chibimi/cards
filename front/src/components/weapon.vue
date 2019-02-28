<template>
	<div class="w-100">
		<div class="form-group row">

			<span class="col-2 text-left">Name</span>
			<span class="col-2 text-left">Type</span>
			<span class="col-1 text-left">rng</span>
			<span class="col-1 text-left">pow</span>
			<span class="col-1 text-left">rof</span>
			<span class="col-1 text-left">aoe</span>
			<span class="col-1 text-left">loc</span>
			<span class="col-1 text-left">cnt</span>
			<span class="col-1 text-left"></span>
			<span class="col-1 text-left"></span>
			
			<input v-model="weapon.name" type="text" class="form-control col-2" placeholder="Name">
			<select v-model="weapon.type" class="form-control col-2">
					<option value=1>Meele</option>
					<option value=2>Ranged</option>
					<option value=3>Mount</option>
				</select>
			<input v-model="weapon.rng" type="text" class="form-control col-1" placeholder="rng">
			<input v-model="weapon.pow" type="text" class="form-control col-1" placeholder="pow">
			<input v-model="weapon.rof" type="text" class="form-control col-1" placeholder="rof">
			<input v-model="weapon.aoe" type="text" class="form-control col-1" placeholder="aoe">
			<input v-model="weapon.loc" type="text" class="form-control col-1" placeholder="loc">
			<input v-model="weapon.cnt" type="text" class="form-control col-1" placeholder="cnt">

			<div v-if="weapon.id" class="col-1 pl-1 pr-0">
				<button type="submit" class="form-control btn btn-success" @click="save(weapon)">Update</button>
			</div>
			<div v-if="weapon.id" class="col-1 pl-1 pr-0">
				<button type="submit" class="form-control btn btn-danger" @click="remove(weapon)">Delete</button>
			</div>
			<div v-if="!weapon.id" class="col-2 pl-1 pr-0">
				<button type="submit" class="form-control btn btn-primary" @click="save(weapon)">Add</button>
			</div>
		</div>
		<hr>
		<!-- <div v-if="weapon.id">
			<h4>Weapon abilities</h4>
			<div class="row">
				<label class="col-1 col-form-label"></label>
				<div class="col-11">
					<Ability
						v-for="(value,index) in abilities"
						:abilitiesList="abilitiesList"
						v-bind:ability="value"
						:key="value.id"
						v-on:remove="removeAbility(value,index)"
					></Ability>
					<Ability :ability="ability" :abilitiesList="abilitiesList" v-on:add="addAbility"></Ability>
				</div>
			</div>
		</div> -->
	</div>
</template>

<script>
import Ability from "./ability.vue";
export default {
	name: "Weapon",
	props: ["abilitiesList", "weapon"],
	components: {
		Ability
	},
	watch: {
		model: function(newVal, oldVal) {
			if (newVal.id === oldVal.id) {
				return;
			}
			console.log("CHANGED WEAPON", newVal, oldVal);
			// watch it
			this.get(newVal.id);
		}
	},
	created: function() {
		console.log("created", this.weapon);
		this.get(this.weapon.id);
	},
	data() {
		return {
			abilities: [],
			ability: {},
		};
	},
	methods: {
		get: function(weaponID) {
			if (!weaponID) {
				return;
			}
			this.$http
				.get("http://localhost:9901/weapons/" + weaponID + "/abilities")
				.then(function(res) {
					console.log(res);
					this.abilities = res.data;
				});
		},
		save: function(weapon) {
			if (weapon.id == null) {
				weapon.id = 0;
			}
			this.$http
				.put("http://localhost:9901/weapons/" + weapon.id, weapon)
				.then(function(res) {
					console.log(res);
					if (res.status === 201) {
						weapon.id = res.data;
						this.$emit("add", weapon);
					}
				});
		},
		remove: function(weapon) {
			console.log("REMOVE FROM MODEL", weapon.id);
			this.$http
				.delete("http://localhost:9901/weapons/" + weapon.id)
				.then(function(res) {
					if (res.status === 204) {
						this.$emit("remove");
					}
				})
				.catch(function(err) {
					console.log(err);
				});
		},
		removeAbility: function(ability, index) {
			console.log("remove abi", ability, index);
			this.$http
				.delete(
					"http://localhost:9901/weapons/" +
						this.weapon.id +
						"/abilities/" +
						ability.id
				)
				.then(function(res) {
					if (res.status === 204) {
						this.abilities.splice(index, 1);
					}
				});
		},
		addAbility: function(ability) {
			this.$http
				.put(
					"http://localhost:9901/weapons/" +
						this.weapon.id +
						"/abilities/" +
						ability.id +
						"?magical=" +
						ability.magical
				)
				.then(function(res) {
					if (res.status === 200) {
						this.abilities.push(ability);
						this.ability = {};
					}
				});
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.statline input {
	max-width: 4rem;
}
</style>
